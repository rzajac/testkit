package testkit

// TODO
// func Test_ReplaceInFile(t *testing.T) {
// 	// --- Given ---
// 	pth := filepath.Join(t.TempDir(), "test_file.txt")
// 	content := `line1
// 	line2
// 	end
// 	`
// 	err := ioutil.WriteFile(pth, []byte(content), 0777)
// 	require.NoError(t, err)
//
// 	// --- When ---
// 	ReplaceInFile(pth, "line", "test")
//
// 	// --- Then ---
// 	got, err := ioutil.ReadFile(pth)
// 	assert.NoError(t, err)
//
// 	exp := `test1
// 	test2
// 	end
// 	`
// 	assert.Exactly(t, exp, string(got))
// }
//
// TODO
// func Test_ModTime(t *testing.T) {
// 	// --- Given ---
// 	fil := TempFile(t.TempDir(), "")
//
// 	// --- When ---
// 	mt := ModTime(fil.Name())
//
// 	// --- Then ---
// 	assert.True(t, time.Now().Sub(mt) < 10*time.Millisecond)
// }
