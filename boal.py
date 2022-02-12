from curses.ascii import NUL


class Solution(object):
    def hasPath(self, maze, start, destination):
        """
        :type maze: List[List[int]]
        :type start: List[int]
        :type destination: List[int]
        :rtype: bool
        """
        row = len(maze)
        col = len(maze[0])
       
        dfs = None
        directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        grid = maze
        def dfsfunc(x, y , src, dst, direction): 
            bit = 0x01<<(direction+1) # b10,b100,b1000,b10000分别是四个方向
            v = grid[x][y]
            # print(v, bit, v&0x01!=1, v&bit==1)
            if v&0x01==1 or v&bit==1: #1 位是1， 或在这个方向走过
                return False
            
            grid[x][y] = v | bit # 标记在这个方向走过
            nx, ny = x+directions[direction][0], y+directions[direction][1]
            if nx < 0 or ny < 0 or nx >= row or ny >= col or grid[nx][ny] == 1:
                # 停球，检查是不是目的地
                # print(x,y, nx,ny)
                if x==destination[0] and y==destination[1]:
                    return True
                else:# 尝试向其他方向走
                    for dirID, dir in enumerate(directions):
                        nx, ny = x+dir[0], y+dir[1]
                        if nx < 0 or ny < 0 or nx >= row \
                            or ny >= col or grid[nx][ny] != src\
                            or dir == direction or abs(dirID-direction)==2: # 同方向或放方向放弃
                            continue
                        if dfs(nx, ny, src, dst, dirID):
                            return True
            elif dfs(nx, ny, src, dst, direction):
                return True
            return False
        dfs = dfsfunc
        for dir in range (0,4):
            if dfs(start[0], start[1], 0, 0, dir):
                return True
        return False
s = Solution()
#maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]]
#start = [0,4]
#destination = [4,4]
maze = [[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]]
start = [4,3]
destination = [0,1]

out = s.hasPath(maze, start, destination)
print(out)