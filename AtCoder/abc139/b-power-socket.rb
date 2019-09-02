a, b = gets.split.map(&:to_i)

num = 0
cnt = 0

num += a
cnt += 1

while num < b
  num += a - 1
  cnt += 1
end

if b <= 1
  puts 0
else
  puts cnt
end
