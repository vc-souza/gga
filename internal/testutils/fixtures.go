package testutils

// CLRS 3rd Edition, 22.2-1
const BasicDG = `
digraph
1#2,4
2#5
3#5,6
4#2
5#4
6#6
`

// CLRS 3rd Edition, 22.2-2
const BasicUG = `
graph
r#s,v
s#r,w
t#u,w,x
u#t,x,y
v#r
w#s,t,x
x#t,u,w,y
y#u,x
`
