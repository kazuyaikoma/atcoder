s = gets.to_s.chomp.split('')

flag = true
odd_group = ['R', 'U', 'D']
even_group = ['L', 'U', 'D']

s.each_with_index do |n, i|
  flag = odd_group.include?(n) if i.even?
  flag = even_group.include?(n) if i.odd?
  break if !flag
end

puts (flag) ? 'Yes' : 'No'

