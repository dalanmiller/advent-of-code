import re

EXAMPLE = """Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II"""

lines = [re.split("[\\s=;,]+", x) for x in EXAMPLE.splitlines()]

# Tunnel Map Dictionary
G = {x[1]: set(x[10:]) for x in lines}

# Valve Pressure Dictionary (skipping 0s)
F = {x[1]: int(x[5]) for x in lines if int(x[5]) != 0}

# Bitfield dict for all the valves that actually have flow rates
# key is the valve, value is the corresponding power-of-two value
I = {x: 1 << i for i, x in enumerate(F)}

# Map of all tunnels as keys with values as
# second map with keys of tunnels with weights.
# . 1 if connected, inf if not
T = {x: {y: 1 if y in G[x] else float("+inf") for y in G} for x in G}

# This then uses Henry-Warshall to compute shortest path
# from each node to every other node.
for k in T:
    for i in T:
        for j in T:
            T[i][j] = min(T[i][j], T[i][k] + T[k][j])


def visit(
    v,  # current valve?
    budget,  # minutes left
    state,  # 15-bit number representing the set of currently open valves
    # I[u] stores a unique power-of-two for each valve with a positive flow (1<<i).
    # The number of such valves in all inputs seems to be limited to 15.
    # So state will a 15-bit number representing the set of currently open valves,
    # I check if a valve is open by testing state&I[u] and set a valve as open with state|I[u].
    # Then, for each distinct state (which represents a distinct set of open valves),
    # I keep the maximum flow achieved in all routes that result in the same set of open valves.
    flow,  # current pressure
    answer,  # dict
):
    answer[state] = max(answer.get(state, 0), flow)

    # Iterate through all the valves because we are going to move to
    # all of them starting from the starting point
    for valve in F:
        # New budget is old budget minus 1 and minus triptime
        newbudget = budget - T[v][valve] - 1

        # Time is out or we've already seen this or this valve is already open
        if I[valve] & state or newbudget <= 0:
            continue

        # Visit next valve
        visit(
            valve,  # valve from loop
            newbudget,  # new budget which is current - 1 and minus travel time
            state
            | I[
                valve
            ],  # bitwise OR which takes current state of valves and adds this valve to on
            flow
            + newbudget
            * F[valve],  # current flow pressure + new flow rate * remaining time
            # (because it's now on to the very end!)
            answer,  # the shared state that is pumped through each recursive iteration
        )

    # bubble up the answer dict which will map 15bit number state to max seen pressure for that state
    return answer


total1 = max(visit("AA", 30, 0, 0, {}).values())
visited2 = visit("AA", 26, 0, 0, {})
total2 = max(
    v1 + v2 for k1, v1 in visited2.items() for k2, v2 in visited2.items() if not k1 & k2
)
print(total1, total2)
