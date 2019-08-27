strs = gets.to_s.split("")

cnt = 0

idx = 1
prev_idx = 0
prev_str = strs[0]

while idx < strs.size
  current_str = strs.slice(idx, idx - prev_idx).join
  idx += 1
  next if prev_str == current_str
  prev_idx = idx - 1
  prev_str = current_str
  cnt += 1
end
puts cnt
