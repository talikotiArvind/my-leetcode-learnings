
rows = 5

def build_pascal_triangle(n):
    output = [None] * n
    for i in range(n):
        row, mid = [1] * (i+1), (i >> 1) + 1
        for j in range(1, mid):
            value = output[i-1][j-1] + output[i-1][j]
            row[j], row[len(row)-j-1] = value, value
        output[i] = row
    print(output)
    return output

build_pascal_triangle(rows)