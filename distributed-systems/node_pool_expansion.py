"""
Node state and status in cassandra:

    Status - U (up) or D (down)
    Indicates whether the node is functioning or not.

    State - N (normal), L (leaving), J (joining), M (moving)
    The state of the node in relation to the cluster.
"""

import time
import random
from enum import Enum
from dataclasses import dataclass


class State(Enum):
    """ Do I need this many states?"""

    NO_STATE = 1
    NORMAL = 2
    JOINING = 3


@dataclass
class Digest:

    state: State
    version: int


class Typology:
    def __init__(self, node):
        self.owner = node.name
        self._cluster = {node.name: Digest(node.state, 0)}

    def node_in_state(self, state: State) -> bool:
        return any(
            [
                digest.state == state
                for name, digest in self._cluster.items()
                if name != self.owner
            ]
        )
        # return state in [digest.state for digest in self._cluster.values()]

    def add(self, name, digest: Digest):
        self._cluster.update({name: digest})

    def __len__(self):
        return len(self._cluster)

    def bump_version(self, state):
        self._cluster[self.owner].version += 1
        self._cluster[self.owner].state = state

    def get_random(self):
        return random.choice(list(self._cluster))

    def items(self):
        for k, v in self._cluster.items():
            yield (k, v)

    def update(self, other):
        for name, digest in other.items():
            if name in self._cluster:
                d = self._cluster[name]
                self._cluster[name] = digest if digest.version > d.version else d
            elif name != self.owner:
                self._cluster[name] = digest


class Node:
    """ Some functions should only be used by the seed node.
    Nodes which are not seeds should not be able to receive join
    requests.
        (Forward join requests to seed nodes?)
    """

    def __init__(self, name, seed=False):
        self.name = name
        self.cluster = "node-pool-e3a1"
        self._state = State.NO_STATE if not seed else State.NORMAL
        self.seed = seed
        self.typology = Typology(self)

        self.current_nodes = len(self.typology)
        self.joining_nodes = 0
        self.adding = []

    @property
    def state(self):
        return self._state

    def _start_orchestrating(self):
        fail_states = [State.JOINING]
        if any([self.typology.node_in_state(state) for state in fail_states]):
            raise Exception("Cluster is rebalancing")

        if self._state != State.JOINING:
            self.current_nodes = len(self.typology)
            self._state = State.JOINING
            self.typology.bump_version(self._state)

    def add(self, node):
        """ Doesn't gossip pending updates until it receives 'done'"""
        if self.cluster != node.cluster:
            raise Exception("Can't add node since it's configured for another cluster")
        self._start_orchestrating()
        digest = Digest(node.state, 0)
        self.adding.append((node.name, digest))
        self.joining_nodes += 1

    def done(self):
        print("Creating typology")
        while self.adding:
            # Time adds artificial delay to enact partition assignment
            time.sleep(0.2)
            name, digest = self.adding.pop()
            self.typology.add(name, digest)
            print("Node added")

        self._state = State.NORMAL
        self.typology.bump_version(self._state)

    def changes_made(self):
        """ Check if all nodes have been added to the system out of those we provided
            partitions for.

            In practice we will also have nodes leaving the system.
        """
        return (len(self.typology) - self.current_nodes) == self.joining_nodes

    def perform(self, network):
        """ Perform gossip protocol.

        https://docs.datastax.com/en/archived/cassandra/3.0/cassandra/architecture/archGossipAbout.html
        """
        # Choose random known node
        node_address = self.typology.get_random()
        # "Contact Node Over Network"
        node = network[node_address]

        # Perform Gossip
        self.gossip(node.typology)
        node.gossip(self.typology)

    def state_check(self):
        if (self.seed) and (self.changes_made()):
            self.joining_nodes = 0
            self.current_nodes = len(self.typology)
            self._state = State.NORMAL
            self.typology.bump_version(self._state)

        # Only once it's joined the cluster can is realise it's
        # a seed itself
        if self._state == State.JOINING and not self.seed:
            print(f"Rebalancing {self.name}")
            time.sleep(0.4)
            self._state = State.NORMAL
            self.typology.bump_version(self._state)

        if self._state == State.NO_STATE:
            self._state = State.JOINING
            self.typology.bump_version(self._state)

    def gossip(self, typology):
        self.state_check()
        self.typology.update(typology)


def print_states(network):
    print([(n.name, n.state) for n in network.values()])


def all_perform(network):
    for n in network.values():
        n.perform(network)
    print_states(network)


if __name__ == "__main__":
    a = Node("A", seed=True)

    b = Node("B")
    c = Node("C")
    f = Node("F")
    network = {"A": a, "B": b, "C": c}  # , "F": f}

    # Add b and c to cluster
    print(a.typology._cluster)
    a.add(b)
    print(a.typology._cluster)
    a.add(c)
    a.done()
    # a.add(f)

    # Perform 4 ticks
    print_states(network)
    all_perform(network)
    # Completed

    print("Adding 'D'")
    d = Node("D")
    network.update({"D": d})
    a.add(d)

    print("Adding 'E'")
    e = Node("E")
    network.update({"E": e})
    a.add(e)

    a.done()

    # Perform 4 ticks
    print_states(network)
    all_perform(network)

    print(a.typology._cluster)
    all_perform(network)
