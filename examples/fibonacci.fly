def "fibonacci" value
  (n: Number)
    fn
      or
        when n == 1 then 1 end
        when n == 2 then 1 end
        when n > 2 then
          def "i" value
            .fibonacci
              n - 2
            end
          end

          def "j" value
            .fibonacci
              n - 1
            end
          end

          i + j
        end
      end
    end
end

array.for
  array 1 2 3 4 5 6 7 8 9 10 end
  
  (e i)
    fn
      std.write
        .fibonacci
          e
        end
      end
    end
end