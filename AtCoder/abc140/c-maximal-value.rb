n = gets.to_i
bs = gets.to_s.split(' ').map(&:to_i)

ans = []
ans[0] = bs[0]
bs.each_with_index do |b, i|
  prev = ans[i]
  idx = i + 1
  ans[idx] = b
  ans[i] = b if prev > b
end

puts ans.inject(:+)
