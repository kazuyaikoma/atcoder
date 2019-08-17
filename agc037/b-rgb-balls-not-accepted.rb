n = gets.to_i
strs = gets.to_s.split('')

rs = strs.map.with_index { |e, i| e == 'R' ? i : nil }.compact.sort!
gs = strs.map.with_index { |e, i| e == 'G' ? i : nil }.compact.sort!
bs = strs.map.with_index { |e, i| e == 'B' ? i : nil }.compact.sort!

copied_rs = rs
copied_gs = gs
copied_bs = bs
border = 0

n.times do
  firsts = [rs.first, gs.first, bs.first]
  lasts = [rs.last, gs.last, bs.last]
  if firsts.max - firsts.min < lasts.max - lasts.min
    copied_rs -= firsts
    copied_gs -= firsts
    copied_bs -= firsts
    border += firsts.max - firsts.min
  else
    copied_rs -= lasts
    copied_gs -= lasts
    copied_bs -= lasts
    border += lasts.max - lasts.min
  end
end

cnt = 0

# TODO: それぞれの組み合わせがborderを満たすかどうか判定したい(product使うとか？)
puts cnt
