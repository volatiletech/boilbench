#!/usr/bin/env ruby

require 'fileutils'

KINDS = %w(Insert Update Delete SelectAll SelectSubset SelectComplex RawBind)

kind_data = {}

KINDS.each { |k| kind_data[k] = {} }

ARGV.each do |file|
  output = []
  File.readlines(file).each do |line|
    KINDS.each do |k|
      next unless line.match(/^[^\s]*#{k}[^\s]*/)
      name, _, nsop, bop, aop = line.gsub(/ns\/op|B\/op|allocs\/op/, '').strip.gsub(/\s{2,}/, ',').gsub(/^[^\/]+\//, '').gsub(/-8/, '').split(',')
      kind_data[k][name] = (kind_data[k][name] || []).push [nsop.to_i,bop.to_i,aop.to_i]
    end
  end
end

stats = {}

kind_data.each_pair do |k, orms|
  stats[k] = {}
  orms.each_pair do |orm, values|
    nsop = values.map { |v| v[0] }
    bop = values.map { |v| v[1] }
    aop = values.map { |v| v[2] }
    stats[k][orm] = {
      nsop: [nsop.min, nsop.max, nsop.inject(:+).to_f / nsop.length.to_f, nsop.max - nsop.min],
      bop: [bop.min, bop.max, bop.inject(:+).to_f / bop.length.to_f, bop.max - bop.min],
      aop: [aop.min, aop.max, aop.inject(:+).to_f / aop.length.to_f, aop.max - aop.min]
    }
  end
end

#stats.each_pair do |k, orms|
#  puts k
#  #puts "name,min ns/op,max ns/op,avg ns/op,err ns/op,min B/op,max B/op,avg B/op,err B/op,min allocs/op,max allocs/op,avg allocs/op,err allocs/op"
#  orms.each_pair do |orm, vals|
#    print "#{orm},"
#    puts [
#      vals[:nsop].map(&:to_s).join(','),
#      vals[:bop].map(&:to_s).join(','),
#      vals[:aop].map(&:to_s).join(','),
#    ].join(',')
#  end
#end


FileUtils::mkdir_p('graph_data')
stats.each_pair do |k, orms|
  [:nsop, :bop, :aop].each do |metric|
    File.open("graph_data/#{k}_#{metric.to_s}.csv", 'w') do |f|
      keys = orms.keys.sort
      keys.each do |orm|
        f.print "#{orm},"
        f.puts orms[orm][metric].map(&:to_s).join(',')
      end
    end
  end
end
