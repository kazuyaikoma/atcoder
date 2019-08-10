def search(a, b)
  return 0 if b < a
  return 1 if a == b
  cnt = 0
  (1..4).each do |a_step|
    (1..4).each do |b_step|
      cnt += search(a + a_step, b - b_step)
    end
  end
  cnt
end

stairs = 10
puts search(0, stairs)
