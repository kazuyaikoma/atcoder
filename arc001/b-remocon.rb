input = gets.to_s.split(' ').map(&:to_i)
a = input[0]
b = input[1]

cnt = 0
num = (b - a).abs
num_str = num.to_s.split('')
if num_str.size > 1
  num_str.pop
  cnt += num_str.join.to_i
end
remaining = num.to_s.split('').last.to_i
if remaining == 9 || remaining == 4
  cnt += 2
elsif remaining == 8
  cnt += 3
elsif remaining == 5
  cnt += 1
elsif remaining > 5
  cnt += 1
  cnt += remaining - 5
else
  cnt += remaining
end

puts cnt
