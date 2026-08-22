def tarjan_scc(n, adj):
    """Iterative — avoids Python's recursion limit on deep graphs."""
    index, low = [-1] * n, [0] * n
    on_stack = [False] * n
    stack, sccs, counter = [], [], 0

    for root in range(n):
        if index[root] != -1:
            continue
        work = [(root, 0)]                    # (node, next-child pointer)
        while work:
            v, pi = work[-1]
            if pi == 0:
                index[v] = low[v] = counter
                counter += 1
                stack.append(v)
                on_stack[v] = True

            descended = False
            for i in range(pi, len(adj[v])):
                w = adj[v][i]
                if index[w] == -1:
                    work[-1] = (v, i + 1)     # resume here after w returns
                    work.append((w, 0))
                    descended = True
                    break
                elif on_stack[w]:
                    low[v] = min(low[v], index[w])
            if descended:
                continue

            if low[v] == index[v]:            # v roots a component
                comp = []
                while True:
                    w = stack.pop()
                    on_stack[w] = False
                    comp.append(w)
                    if w == v:
                        break
                sccs.append(comp)

            work.pop()
            if work:
                u = work[-1][0]
                low[u] = min(low[u], low[v])
    return sccs