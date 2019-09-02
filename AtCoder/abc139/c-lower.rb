n = gets.to_i
nums = gets.to_s.split.map(&:to_i)

cnt = 0
prev_cnt = 0
prev = nums.first
nums.each_with_index do |num, idx|
  next if idx == 0
  if prev >= num
    cnt += 1
  else
    prev_cnt = [cnt, prev_cnt].max
    cnt = 0
  end
  prev = num
end

puts [cnt, prev_cnt].max
