require 'date'
require 'benchmark'

puts Benchmark::CAPTION
puts Benchmark.measure {
  from_left = 1964_10_10.to_s(2)[4..11].to_i(2)   # => 93
  to_left   = 2020_07_24.to_s(2)[4..11].to_i(2)   # => 161

  (from_left..to_left).each do |n|
    left  = '%08b' % n     # %フォーマットにより、0埋め8ケタの2進数に変換
    right = left.reverse

    (0..1).each do |m|
      date = "1001#{left}#{m}#{right}1001"
      begin
        p Date.parse(date.to_i(2).to_s).strftime('%Y年%-m月%-d日')
      rescue
      end
    end
  end
}
