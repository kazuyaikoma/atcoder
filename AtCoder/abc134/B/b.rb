n, d = gets.split.map(&:to_i)
count = 0
while n > 0 do
  n -= 2 * d + 1
  count += 1
end

puts count