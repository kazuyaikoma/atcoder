n = gets.to_i
array = []
n.times { array.push(gets.to_i) }

m = array.max
cpd = array.clone
cpd.delete(m)
second_m = cpd.max
second_m = m if array.count(m) > 1

array.each do |a|
  if a < m
    puts m
  elsif 
    puts second_m
  end
end
