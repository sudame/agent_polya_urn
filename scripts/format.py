import pandas as pd
import numpy as np

nodes = pd.read_table("agents")
nodes.columns = ["id"]
nodes["label"] = nodes["id"]
nodes.to_csv("nodes.csv", index=False)


edges = pd.DataFrame(data=np.loadtxt("event", dtype=int), columns=["source", "target"])
edges["weight"] = 1
edges.groupby(
    by=[
        "source",
        "target",
    ]
).sum().reset_index().to_csv("edges.csv", index=False)
