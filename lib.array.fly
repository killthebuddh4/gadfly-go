def "slice" value
  (arr start finish)
    fn 
      array.filter arr 
        (el i)
          fn
            def "result" value
              and
                when i >= start then true end
                when i < finish then true end
              end
            end

            result == true
          end
      end
    end
end

def "length" value
  (arr)
    fn
      def "result" value 0 end

      array.for arr
        (e i)
          fn
            let "result" value i end
          end
      end

      result + 1
    end
end

def "includes" value
  (arr e)
    fn
      def "includes" value false end

      array.for arr
        (test i)
          fn
            when test == e then
              let "includes" value true end
            end
          end
      end

      includes
    end
end

def "equal" value
  (arr_a arr_b)
    fn
      def "result" value true end

      def "arr_a_length" value
        .length arr_a end
      end

      def "arr_b_length" value
        .length arr_b end
      end

      when arr_a_length != arr_b_length then
        let "result" value false end
      end

      array.for arr_a
        (e i)
          fn
            def "b" value
              array.read arr_b i end
            end

            when e != b then
              let "result" value false end
            end
          end
      end

      result
    end
end