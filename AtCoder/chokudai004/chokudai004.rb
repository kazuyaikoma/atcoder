n, b1, b2, b3 = gets.split.map(&:to_i)

double_l_array = []
double_r_array = []
n.times { double_l_array << gets.split.map(&:to_i) }
n.times { double_r_array << gets.split.map(&:to_i) }

(0...n).each do |i|
  double_l_array[i].each_with_index do |num, idx|
    print num
    print ' ' if idx != n - 1
  end
  puts '' if i != n - 1
end
