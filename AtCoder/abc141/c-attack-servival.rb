n, k, q = gets.split.map(&:to_i)

arr = Array.new(n, k)
arr = arr.map { |a| a - q }

q.times { arr[gets.to_i - 1] += 1 }

arr.each do |a|
  puts (a > 0) ? 'Yes' : 'No'
end
