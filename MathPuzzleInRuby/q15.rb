@steps = 10
@jump = 4
@cnt = 0

def search_patterns(a_step, b_step)
  return if a_step > b_step
  (1..@jump).each do |a|
    (1..@jump).each do |b|
      a_pos = a_step + a
      b_pos = b_step - b
      if a_pos == b_pos
        @cnt += 1
      else
        search_patterns(a_pos, b_pos)
      end
    end
  end
end

search_patterns(0, @steps)

puts @cnt
