n = gets.to_i
hs = gets.split(' ').map(&:to_i)
dp = Array.new(n, 0)
dp[1] = (hs[1] - hs[0]).abs

(2...n).each do |i|
  dp[i] = [
    dp[i - 1] + (hs[i] - hs[i - 1]).abs,
    dp[i - 2] + (hs[i] - hs[i - 2]).abs
  ].min
end

p dp[n - 1]
