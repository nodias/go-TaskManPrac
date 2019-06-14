package task

// func Test_status_MarshalJSON(t *testing.T) {
// 	s := TODO
// 	b, err := s.MarshalJSON()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if string(b) == "TODO" {
// 		fmt.Println("success")
// 	}
// }

// func Test_status_UnmarshalJSON(t *testing.T) {
// 	s_string := "TODO"
// 	var s status
// 	err := s.UnmarshalJSON([]byte(s_string))
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if s == TODO {
// 		fmt.Println("success")
// 	}
// }

// func Test_Deadline_MarshalJSON(t *testing.T) {
// 	d := time.Now().Add(time.Hour * 5)
// 	str_d, err := d.MarshalJSON()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if string(str_d) == "1560327633" {
// 		fmt.Println("success")
// 	}
// 	fmt.Println(string(str_d))
// }

// func Test_Deadline_UnmarshalJSON(t *testing.T) {
// 	deadline := NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local))
// 	str_deadline, _ := deadline.MarshalJSON()

// 	var d Deadline
// 	data := []byte(str_deadline)
// 	err := d.UnmarshalJSON(data)
// 	if err != nil {
// 		t.Errorf("Test_Deadline_UnmarshalJSON error : %s", err)
// 		return
// 	}

// 	if d.Time == deadline.Time {
// 		fmt.Println("success")
// 		return
// 	}
// 	t.Errorf("Test_Deadline_UnmarshalJSON error : \n expected '%v' \n but '%v'", deadline, d.Time)
// }

// func ExampleTaskMarshal() {
// 	//Marshal
// 	t := Task{
// 		"laundry",
// 		TODO,
// 		NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 		[]Task{
// 			Task{
// 				"put",
// 				TODO,
// 				NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 				nil,
// 			},
// 		},
// 	}
// 	b, err := json.Marshal(t)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(b))
// 	//Unmarshal
// 	t2 := Task{}
// 	err = json.Unmarshal(b, &t2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if !reflect.DeepEqual(t, t2) {
// 		resultstr := fmt.Sprintf("\n expected %v \n but %v", t, t2)
// 		fmt.Println(errors.New(resultstr))
// 	}

// 	//Output:
// 	//{"Title":"laundry","Status":"TODO","Deadline":1560351600,"SubTask":[{"Title":"put","Status":"TODO","Deadline":1560351600,"SubTask":null}]}
// }

// func ExampleTest() {
// 	//Marshal
// 	t := Task{
// 		"laundry",
// 		TODO,
// 		NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 		[]Task{
// 			Task{
// 				"put",
// 				DONE,
// 				NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 				nil,
// 			},
// 			Task{
// 				"wash",
// 				TODO,
// 				NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 				[]Task{
// 					Task{
// 						"put Detergent",
// 						TODO,
// 						NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 						nil,
// 					},
// 					Task{
// 						"start washer",
// 						TODO,
// 						NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 						nil,
// 					},
// 				},
// 			},
// 			Task{
// 				"take out",
// 				TODO,
// 				NewDeadline(time.Date(2019, 06, 13, 00, 00, 00, 00, time.Local)),
// 				nil,
// 			},
// 		},
// 	}
// 	fmt.Println(IncludeSubTasks(t))

// 	//Output:
// }
