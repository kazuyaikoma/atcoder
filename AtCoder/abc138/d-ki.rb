# This is NOT COMPLETED (got TLE) code
first_input = gets.to_s.split(' ').map(&:to_i)
n = first_input.first
q = first_input.last

# create node array
nodes = []
(1..n).each { |node| nodes << node }

(n - 1).times do
  input = gets.to_s.split(' ').map(&:to_i)
  idx = nodes.index(input.first) + 1
  nodes.delete(input.last)
  nodes.insert(idx, input.last)
end

values = Array.new(n, 0)
q.times do
  input = gets.to_s.split(' ').map(&:to_i)
  idx = nodes.index(input.first)
  sliced_nodes = nodes.slice(idx..n - 1)
  sliced_nodes.each do |i|
    values[i - 1] += input.last
  end
end

values.each do |value|
  print value
  print ' '
end
puts ''
