def "status" value nil end
def "got_sig_test_a" value false end
def "got_sig_test_b" value false end

signal
  SIG_TEST_A

  (signal)
    fn
      let "got_sig_test_a" value true end
      let "got_sig_test_b" value false end
      FEED_OK
    end
end

signal
  SIG_TEST_B

  (signal)
    fn
      let "got_sig_test_a" value false end
      let "got_sig_test_b" value true end
      FEED_KILL
    end
end

emit
  SIG_TEST_A

  on
    FEED_OK

    (feedback)
      fn
        let "status" value "WE ARE OK" end
      end
  end

  on
    FEED_KILL

    (feedback)
      fn
        let "status" value "OH NO, WE'RE DED" end
      end
  end
end

if
  status != "WE ARE OK"
then
  std.write "FAIL: Ok status" end
else
  std.write "PASS: Ok status" end
end

if
  got_sig_test_a
then
  std.write "PASS: Got SIG_TEST_A" end
else
  std.write "FAIL: Got SIG_TEST_A" end
end

if
  got_sig_test_b
then
  std.write "FAIL: Got SIG_TEST_B" end
else
  std.write "PASS: Got SIG_TEST_B" end
end

emit
  SIG_TEST_B

  on
    FEED_OK

    (feedback)
      fn
        let "status" value "WE ARE GREAT" end
      end
  end

  on
    FEED_KILL

    (feedback)
      fn
        let "status" value "OH NO, WE'RE DED" end
      end
  end
end

if
  status != "OH NO, WE'RE DED"
then
  std.write "FAIL: Ok sttatus" end
else
  std.write "PASS: Ok status" end
end

if
  got_sig_test_a
then
  std.write "FAIL: Got SIG_TEST_A" end
else
  std.write "PASS: Got SIG_TEST_A" end
end

if
  got_sig_test_b
then
  std.write "PASS: Got SIG_TEST_B" end
else
  std.write "FAIL: Got SIG_TEST_B" end
end

