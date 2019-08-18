n = gets.to_i
values = gets.to_s.split(' ').map(&:to_f).sort!

prev_value = (values[0] + values[1]) / 2

if n > 2
  (2...n).each do |idx|
    prev_value = (prev_value + values[idx]) / 2
  end
end

puts prev_value
