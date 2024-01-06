def slice
  fn |arr, start, finish|
    filter arr 
      fn |el, i|
        def result
          and
            i >= start true
            i < finish true
          end
        end

        call print result end

        result != false
      end
    end
  end
end

def length
  fn |arr|
    def result 0 end

    for arr
      fn |e, i|
        call print "hello" end
        let result i end
        call print result end
      end
    end

    call print "final" end
    result
  end
end