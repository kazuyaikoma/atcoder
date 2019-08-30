read = 'READ'
write = 'WRITE'
talk = 'TALK'
skill = 'SKILL'
# first characters of each words
heads = "#{read[0]}#{write[0]}#{talk[0]}#{skill[0]}"

cnt = 0

all_patterns = "#{read}#{write}#{talk}#{skill}".split('').uniq.join
[*0..9].permutation.each do |nums|
  is_break = false
  sentence = "#{read}+#{write}+#{talk}==#{skill}"
  (0..9).each do |i|
    if nums[i].zero? && heads.include?(all_patterns[i])
      is_break = true
      break
    end
    sentence.gsub!(/#{all_patterns[i]}/, nums[i].to_s)
  end
  next if is_break
  cnt += 1 if eval(sentence)
  p sentence if eval(sentence)
end

p cnt
