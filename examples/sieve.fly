# An innefficient sieve of Eratosthenes.

def "is_prime" value
  (n)
    fn
      def "composites" value
        array end
      end

      def "base" value 2 end
      
      def "test" value base end

      while base < n
        do
          let "test" value base end

          while test < n
            do
              let "test" value test + base end


              let "composites" value
                array.push composites
                  test
                end
              end
            end
          end

          let "base" value base + 1 end
        end
      end

      def "is_composite" value
        .includes composites n
        end
      end

      is_composite == false
    end
end


def "k" value 0 end

while k < 100
  do
    when .is_prime k end then
      std.write k end
    end

    let "k" value k + 1 end
  end
end