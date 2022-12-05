# effect-date

Very simple go utility for getting an effect date (future or past) from an input isodate (yyyy-mm-dd) taking care about weekends and **french** holidays.

The holidays calendar is filled 'til 2027-12-31.

## installation

```
go get github.com/darthyoh/effect-date
```

## usage

```
package main

import "github.com/darthyoh/effect-date"

func main() {
    //simple usage
    result, _ := GetEffectDate("2022-12-05", 2, false) // => 2022-12-07

    //invalid input isodates return errors
    _, err := GetEffectDate("20220-12-05", 2, false) // => error !!!

    //if the openDays flag is set to false, GetEffectDate doesn't care about week ends or holidays
    result, _ := GetEffectDate("2022-12-05",12, false) // => 2022-12-17

    //but is openDays is true, GetEffectDate will skip week ends...
    result, _ := GetEffectDate("2022-12-05", 7, true) // => 2022-12-14

    //... or holidays ...
    result, _ := GetEffectDate("2022-07-11", 3, true) // => 2022-07-15

    //... or both !!!
    result, _ := GetEffectDate("2022-07-11", 15, true) // => 2022-08-02

    //If a holiday is a week-end day, it is skipped once
    result, _ := GetEffectDate("2022-12-23",2,true) // => 2022-12-27

    //Delay can be negative (reverse effect date)...
    result, _ := GetEffectDate("2022-12-07",-2,false) // => 2022-12-05

    //...and can take care about holidays and weed-ends !!!
    result, _ := GetEffectDate("2022-08-02",-15,true) // => 2022-07-11
}
```
