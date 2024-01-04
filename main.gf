let x 300

let y 200

let f if x < y then
  fn |x|
    x * 10
  end
else
  fn |x|
    x * 20
  end
end


let times fn |n, f|
  if n == 0 then
    0
  else
    f()
    times(n - 1, f)
  end
end


let z fn ||
  print("hello")
end


times(10, z)

let a fn || "goodbye" end

print(a())