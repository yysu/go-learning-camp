需要使用 Wrap 抛给上层：
- Dao的作用是封装对数据库的访问：增删改查，不涉及业务逻辑，只是达到按某个条件获得指定数据的要求.
- 不确定业务形态， 所以应该将这个error提交给上层。由上层业务场景确定代码逻辑。


dao:
```
return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))
```

biz:
```
if errors.Is(err, code.NotFound} {
  // handle other logic
}
```