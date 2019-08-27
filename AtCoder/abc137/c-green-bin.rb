n = gets.to_i
strs = []
n.times do
  strs << gets.to_s.chomp
end

sorted_strs = []
strs.each do |str|
  sorted_strs << str.chars.sort.join
end

arr = sorted_strs.group_by(&:itself).values.select { |value| value.size > 1 }
cnt = 0
arr.each do |elm|
  if elm.size > 1
    (1..elm.size - 1).each { |num| cnt += num }
  end
end
puts cnt
