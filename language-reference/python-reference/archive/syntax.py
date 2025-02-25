my_fav_color = 'red'
your_fav_color = 'blue'

fav_colors = [] # what a comment!

colors = ['green', 'red', 'yellow', 'blue']

for i in range(len(colors)):
    if (colors[i] == my_fav_color or colors[i] == your_fav_color):
        print(i)

# set the midpoint
midpoint = 5

# semicolons for multi-statement line
lower = []; upper = []

# split the numbers into lower and upper
for i in range(10):
    if (i < midpoint):
        lower.append(i)
    else:
        upper.append(i)
        
print("lower:", lower)
print("upper:", upper)

# multi-line!
x = 1 + 2 + 3 + 4 +\
    5 + 6 + 7 + 8

# also multi-line!
x = (1 + 2 + 3 + 4 +
     5 + 6 + 7 + 8)

