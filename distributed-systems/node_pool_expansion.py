import time
import random
from enum import Enum
from dataclasses import dataclass


class State(Enum):
    """ Do I need this many states?"""

    NORMAL = 1
    PENDING = 2
    ORCHESTRATING = 3
    ORCHESTRATED = 4
    JOINING = 5
    NO_STATE = 6


@dataclass
class Digest:

    state: State
    version: int


class Typology:
    def __init__(self, node):
        self.owner = node.name
        self._cluster = {node.name: Digest(node.state, 0)}

    def node_in_state(self, state: State) -> bool:
        return state in [digest.state for digest in self._cluster.values()]

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
        for node, digest in other.items():
            if node in self._cluster:
                d = self._cluster[node]
                self._cluster[node] = digest if digest.version > d.version else d
            else:
                self._cluster[node] = digest


class Node:
    """ Some functions should only be used by the seed node.
    Nodes which are not seeds should not be able to receive join
    requests.
        (Forward join requests to seed nodes?)
    """

    def __init__(self, name):
        self.name = name
        self.cluster = "node-pool-e3a1"
        self._state = State.NO_STATE
        self.typology = Typology(self)

        self.current_nodes = len(self.typology)
        self.joining_nodes = 0
        self.adding = []

    @property
    def state(self):
        return self._state

    def _start_orchestrating(self):
        fail_states = [State.ORCHESTRATED, State.JOINING]
        if any([self.typology.node_in_state(state) for state in fail_states]):
            raise Exception("Nodes already being added")

        if self._state != State.ORCHESTRATING:
            self.current_nodes = len(self.typology)
            self._state = State.ORCHESTRATING
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

        self._state = State.ORCHESTRATED
        self.typology.bump_version(self._state)

    def check_added(self):
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

    def gossip(self, typology):
        if self._state == State.JOINING:
            print(f"Rebalancing {self.name}")
            time.sleep(0.4)
            self._state = State.NORMAL
            self.typology.bump_version(self._state)

        if self._state == State.NO_STATE:
            self._state = State.JOINING
            self.typology.bump_version(self._state)

        if (self._state == State.ORCHESTRATED) and (self.check_added()):
            self.joining_nodes = 0
            self.current_nodes = len(self.typology)
            self._state = State.NORMAL
            self.typology.bump_version(self._state)

        self.typology.update(typology)


def create_cluster():
    n = Node("A")
    n._state = State.NORMAL
    return n


def print_states(network):
    print([(n.name, n.state) for n in network.values()])


def all_perform(network):
    for n in network.values():
        n.perform(network)
    print_states(network)


if __name__ == "__main__":
    a = create_cluster()
    b = Node("B")
    c = Node("C")
    f = Node("F")
    network = {"A": a, "B": b, "C": c}  # , "F": f}

    # Add b and c to cluster
    a.add(b)
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
