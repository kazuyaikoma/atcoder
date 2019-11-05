n, k = gets.split.map(&:to_i)
hs = gets.split.map(&:to_i)
dp = Array.new(n, 0)
dp[1] = (hs[1] - hs[0]).abs

# i歩目時点でのdp値を計算
(2...n).each do |i|
  temp = []
  # j歩目から計算
  min_j = i - k
  min_j = 0 if min_j.negative?
  (min_j...i).each do |j|
    t = dp[j] + (hs[i] - hs[j]).abs
    temp.push(t)
  end
  dp[i] = temp.min
end

p dp[n - 1]
