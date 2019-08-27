n = gets.to_i
as = gets.to_s.split(' ').map(&:to_f)

nums = []
as.each do |a|
  nums << 1 / a
end
num = nums.inject(:+)

puts 1 / num
