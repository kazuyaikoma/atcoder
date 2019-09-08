n, k = gets.split.map(&:to_i)
s = gets.to_s.split('')

score = 0
s.each_cons(2) { |s1, s2| score += 1 if s1 == s2 }
puts [ 2 * k + score, n  - 1 ].min
