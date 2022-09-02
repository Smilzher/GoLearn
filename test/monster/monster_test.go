package monster

import "testing"

//测试用例
func TestStore(t *testing.T) {
	//先创建一个Monster实例
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望为=%v 实际为=%v", true, false)
	}
	t.Logf("monster.Store() 测试成功！")
}

func TestReStore(t *testing.T) {
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望为=%v, 实际为=%v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore()错误，希望为=%v 实际为=%v", "红孩儿", monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功")
}
