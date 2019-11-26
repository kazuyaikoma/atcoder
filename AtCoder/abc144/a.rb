a, b = gets.chomp.split.map(&:to_i)
if 1 <= a && a <= 9 && 1 <= b && b <= 9
  puts a * b
else
  puts -1
end
