require 'pry'
require 'flog'

# if nil romnum has pattern
# if not nil no pattern
class Romannum
  ROMANS_ONES = {
    '1': 'i', '2': nil, '3': nil,
    '4': 'iv',
    '5': 'v', '6': nil, '7': nil, '8': nil,
    '9': 'ix'
  }

  ROMANS_TENS = {
    '1': 'x', '2': nil, '3': nil,
    '4': 'xl',
    '5': 'l', '6': nil, '7': nil, '8': nil,
    '9': 'xc'
  }

  ROMANS_HUNDREADS = {
    '1': 'c', '2': nil, '3': nil,
    '4': 'cd',
    '5': 'd', '6': nil, '7': nil, '8': nil,
    '9': 'cm'
  }

  ARR_INDEX_TO_KETA_NUM = {
    '0': ROMANS_HUNDREADS,
    '1': ROMANS_TENS,
    '2': ROMANS_ONES
  }

  attr_accessor :row_num_ar, :roman

  def initialize(row_num)
    @row_num_ar = sprit_nums(row_num)
    @roman = call_roman.join('')
  end

  def sprit_nums(num)
    num.to_s.split('').map(&:to_i)
  end

  def call_roman
    row_num_ar.map.with_index do |row_num, keta|
      return_roman(row_num, keta)
    end
  end

  def return_roman(row_num, keta)
    if ARR_INDEX_TO_KETA_NUM[:"#{keta}"][:"#{row_num}"] == nil?
      num_converter(row_num, ARR_INDEX_TO_KETA_NUM[:"#{keta}"])
    else
      ARR_INDEX_TO_KETA_NUM[:"#{keta}"][:"#{row_num}"]
    end
  end

  def num_converter(num, roman_hash)
    if num.between?(1, 3)
      roman_hash[:'1'] + roman_hash[:'1'] * (num - 1)
    elsif num.between?(6, 8)
      roman_hash[:'5'] + roman_hash[:'4'][0] * (num - 5)
    end
  end
end

p Romannum.new(432).roman == 'cdxxxii'
p Romannum.new(632).roman
p Romannum.new(999).roman == 'cmxcix'
p Romannum.new(876).roman
p Romannum.new(923).roman == 'cmxxiii'
