n1 = gets.to_s.split('')
n2 = gets.to_s.split('')

cnt = -1
n1.each_with_index do |n, i|
  cnt += 1 if n == n2[i]
end

puts cnt
