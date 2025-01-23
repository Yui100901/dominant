package geo

import (
	"github.com/Yui100901/MyGo/log_utils"
	"math"
)

//
// @Author yfy2001
// @Date 2024/8/23 19 08
//

//部分地理相关计算函数

// EarthRadius 地球平均半径
const EarthRadius = 6371_393

// ExecOffset 传入点p
// 方向角正北为0（以度为单位）
// 移动的距离（以米为单位）
func ExecOffset(p *Coordinate, cf *CoordinateOffset) *Coordinate {
	log_utils.Info.Println("当前位置：", p.Longitude, p.Latitude)

	newLongitude := p.Longitude + cf.LongitudeOffset
	newLatitude := p.Latitude + cf.LatitudeOffset
	log_utils.Info.Println("偏移位置：", newLongitude, newLatitude)
	return NewCoordinate(newLongitude, newLatitude)
}

// CalcCoordinateOffset 计算坐标偏移
func CalcCoordinateOffset(c *Coordinate, of *AzimuthOffset) *CoordinateOffset {
	// 方向角度转换为弧度
	azimuthRadians := DegreeToRadians(of.Azimuth)
	//纬度转换成弧度
	latitudeRadians := DegreeToRadians(c.Latitude)
	//单位经度距离
	unitLongitudeDistance := 1 * 2 * math.Pi * math.Cos(latitudeRadians) * EarthRadius / 360
	//单位纬度距离
	unitLatitudeDistance := 1 * 2 * math.Pi * EarthRadius / 360
	//单位距离的经度变化量
	unitDeltaLongitude := 1 * math.Sin(azimuthRadians) / unitLongitudeDistance
	//单位距离经度变化量
	unitDeltaLatitude := 1 * math.Cos(azimuthRadians) / unitLatitudeDistance

	//经度偏移
	deltaLongitude := unitDeltaLongitude * of.Distance
	//纬度偏移
	deltaLatitude := unitDeltaLatitude * of.Distance

	return NewCoordinateOffset(deltaLongitude, deltaLatitude)

}

// DegreeToRadians 角度转弧度
func DegreeToRadians(degree float64) float64 {
	return degree * math.Pi / 180
}

// RadiansToDegree 弧度转角度
func RadiansToDegree(radians float64) float64 {
	return radians * 180 / math.Pi
}

// Haversine 公式，计算两点间距离
func Haversine(p1, p2 *Coordinate) float64 {

	//经度变化量
	dLongitudeRadians := p2.LongitudeRadians - p1.LongitudeRadians
	//纬度变化量
	dLatitudeRadians := p2.LatitudeRadians - p2.LatitudeRadians

	a := math.Sin(dLatitudeRadians/2)*math.Sin(dLatitudeRadians/2) +
		math.Cos(p1.Latitude*math.Pi/180.0)*
			math.Cos(p2.Latitude*math.Pi/180.0)*
			math.Sin(dLongitudeRadians/2)*math.Sin(dLongitudeRadians/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return EarthRadius * c
}
