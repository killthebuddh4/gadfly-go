def "is_fizz" value
  (n)
    fn
      def "r" value
        .math.mod 
          n 3
        end
      end

      r == 0
    end
end

def "is_buzz" value
  (n)
    fn
      def "r" value
        .math.mod 
          n 5
        end
      end

      r == 0
    end
end

def "is_fizzbuzz" value
  (n)
    fn
      def "both" value
        and
          when .is_fizz n end then true end

          when .is_buzz n end then true end
        end
      end

      both == true
    end
end

def "fizzbuzz" value
  (n)
  fn
    def "k" value 1 end

    while k <= n
      do
        std.write k end

        or
          when .is_fizzbuzz k end then
            std.write "fizzbuzz" end
            true
          end

          when .is_fizz k end then
            std.write "fizz" end
            true
          end

          when .is_buzz k end then
            std.write "buzz" end
            true
          end
        end

        let "k" value k + 1 end
      end
    end
  end
end

.fizzbuzz 20 end


