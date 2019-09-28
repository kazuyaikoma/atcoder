require 'prime'

a, b = gets.split.map(&:to_i)
gcd = a.gcd(b)
cnt = 1

(1..gcd).each do |n|
  if Prime.prime?(n)
    cnt += 1 if (a % n).zero? && (b % n).zero?
  end
end

puts cnt
