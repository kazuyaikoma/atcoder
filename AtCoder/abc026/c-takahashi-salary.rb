Person = Struct.new(:salary, :subs_salaries, :boss_idx)

n = gets.to_i
persons = [Person.new(1, [], -1)]
(1..n - 1).each do |i|
  persons[i] = Person.new(1, [], gets.to_i - 1)
end

(n - 1).downto(0) do |i|
  if persons[i].subs_salaries.empty?
    persons[i].salary = 1
  else
    persons[i].salary = persons[i].subs_salaries.minmax.inject(:+) + 1
  end
  boss_idx = persons[i].boss_idx
  persons[boss_idx].subs_salaries << persons[i].salary
  persons[boss_idx].salary = persons[boss_idx].subs_salaries.minmax.inject(:+) + 1
end

puts persons.first.salary
