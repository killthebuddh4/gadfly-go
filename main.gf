def x 300

def y 200

def f if x < y then
  fn |x|
    x * 10
  end
else
  fn |x|
    x * 20
  end
end


def times fn |n, f|
  if n == 0 then
    0
  else
    f()
    times(n - 1, f)
  end
end


def z fn ||
  print("hello")
end


times(10, z)

def a fn || "goodbye" end

print(a())