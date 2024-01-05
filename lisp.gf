def x 200 end

def y 200 end

def z 
  fn || do
    print("hello")

    def y 
      fn || do
        print(" world")
      end
    end

    y()
  end
end

z()

if
  when x > y end
  then print("x is greater than y") end
  else print("x is not greater than y") end
end

def a 10 end
def b 20 end

def c 0 end
def d 40 end


def k
  or
      a == 100000
      false
  end
end

print(k)


def arr
  array
    "hey"
    "cool"
    "cool"
  end
end

print(arr)

for arr
  fn |e, i| do
    print("test")
    print(e)
    print(i)
  end
end

def t
  map arr
    fn |e, i| do
      i * i
    end
  end
end

def r
  reduce t
    0

    fn |r, e, i| do
      r + e
    end
  end
end


def p
  def n 0 end
  do
    if true then
      r
    end else
      20
    end
    end
  end
end


print(p)