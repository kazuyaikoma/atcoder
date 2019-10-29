N = 12
loc = [0, 0]

def move(path)
  return 1 if path.count == N + 1
  
  cnt = 0
  [[1, 0], [-1, 0], [0, 1], [0, -1]].each do |n|
    x = path.last[0] + n[0]
    y = path.last[1] + n[1]
    new_loc = [x, y]
    cnt += move(path + [new_loc]) unless path.include?(new_loc)
  end
  cnt
end

p move([loc])