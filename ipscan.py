import pandas
from operator import itemgetter
from netaddr import IPNetwork

colnames = ['routes']
data = pandas.read_csv("./data-in.csv", names=colnames)

routes = data.routes.tolist()
def createNestedList(list):
    return [[el] for el in list]

routes = createNestedList(routes)

def splitRoutes(list):
    resultList = []
    for temp in list:
        subList = temp[0].split('/')
        temp.append(subList[0])
        temp.append(subList[1])
    return(list)

routes = splitRoutes(routes)

routes = sorted(routes, key=itemgetter(2))

for i in routes:
    network1 = i[0]
    for j in routes:
        network2 = j[0]
        if (IPNetwork(network1) in IPNetwork(network2)) and (IPNetwork(network1) != IPNetwork(network2)):
            print(network1, " is in ", network2)

