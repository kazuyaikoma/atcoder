n, m = gets.split.map(&:to_i)
stick = 1
cnt = 0

while (n - stick).positive?
  stick += [stick, m].min
  p stick
  cnt += 1
end

p cnt
