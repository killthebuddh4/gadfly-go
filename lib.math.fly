def "math.mod" value
  (n k)
    fn
      or
        when n < k then n end
        when n == k then 0 end
        when n > k then
          .math.mod
            n - k
            k
          end
        end
      end
    end
end

def "math.factorial" value
  (n)
    fn
      or
        when n == 0 then 1 end
        when n > 0 then
          def "inner" value
            .math.factorial
              n - 1
            end
          end

          n * inner
        end
      end
    end
end
