s = gets.to_s.chomp

ans = ""
case s
when "Sunny" then
  ans = "Cloudy"
when "Cloudy" then
  ans = "Rainy"
when "Rainy" then
  ans = "Sunny"
end

puts ans