require 'date'
require 'benchmark'

puts Benchmark::CAPTION
puts Benchmark.measure {
(Date.parse('1964-10-10')..Date.parse('2020-7-24')).each do |date|
  original = date.strftime('%Y%0m%0d').to_i
  mutated = original.to_s(2).reverse.to_i(2)
  p date.strftime('%Y年%-m月%-d日') if original == mutated
end
}
