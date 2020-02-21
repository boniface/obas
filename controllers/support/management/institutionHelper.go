package management

import (
	"fmt"
	"obas/io/academics"
	institutionIO "obas/io/institutions"
	"obas/io/location"
)

type InstitutionHolder struct {
	InstitutionId     string
	InstitutionTypeId string
	Institution       string
	InstitutionType   string
}
type InstitutionLocHolder struct {
	Id          string
	Institution string
	Town        string
	Longitude   string
	Latitude    string
}

func GetInstitutionLocation() []InstitutionLocHolder {
	var institutionLocation []InstitutionLocHolder
	institutsLocation, err := institutionIO.ReadInstitutionLocations()
	if err != nil {
		println("error reading institutionLocations")
	} else {
		for _, institutionLoca := range institutsLocation {
			//fmt.Println("error in reading townNamw in InstitutionManagementHandler method", institutionLoca.LocationId)

			institutionName, errr := institutionIO.GetInstitution(institutionLoca.InstitutionId)
			if errr != nil {
				fmt.Println("error in reading institutionName in InstitutionManagementHandler method")
			}
			townNamw, err := location.GetLocation(institutionLoca.LocationId)
			//fmt.Println("error in reading townNamw in InstitutionManagementHandler method", townNamw)
			if err != nil {
				fmt.Println("error in reading townNamw in InstitutionManagementHandler method")
			}
			institutionLocation = append(institutionLocation, InstitutionLocHolder{institutionLoca.InstitutionId, institutionName.Name, townNamw.Name, institutionLoca.Longitude, institutionLoca.Latitude})
		}
	}
	return institutionLocation
}

/***
this receiver an institution id and returns
an institutionHolder(this object contains institutionId,
institutionName,intitutionType id and institutionType)
****/

func GetInstitutionHolder() []InstitutionHolder {
	var institutionsHolder []InstitutionHolder

	institution, err := institutionIO.GetInstitutions()
	if err != nil {
		fmt.Println("error reading institution in getInstitutionHolder")
	}
	/**Reading all the institution and their type**/

	for _, value := range institution {
		institutionType, err := institutionIO.GetInstitutionType(value.InstitutionTypeId)
		if err != nil {
			fmt.Println("error reading institutionType in getInstitutionHolder")
		} else {
			result := InstitutionHolder{value.Id, institutionType.Id, value.Name, institutionType.Name}
			institutionsHolder = append(institutionsHolder, result)
		}
	}
	return institutionsHolder
}

/**this method returns all the institutions and their Address**/
type InstitutionAddressHolder struct {
	InstitutionAddressId string
	AddressTypeId        string
	Institution          string
	Address              string
	Postal               string
}

func GetInstitutionAddress() []InstitutionAddressHolder {
	var institutionsAddressHolder []InstitutionAddressHolder
	institutionAddresses, err := institutionIO.GetAllInstitutionAddresses()
	if err != nil {
		fmt.Println("error reading institutionAddresses in GetInstitutionAddress")
	} else if institutionAddresses != nil {
		for _, institutionAddrres := range institutionAddresses {
			myInstitution, err := institutionIO.GetInstitution(institutionAddrres.InstitutionId)
			if err != nil {
				fmt.Println("An error in InstitutionManagementHandler when reading myInstitution")

			} else if myInstitution.Name != "" {
				institutionAddress := InstitutionAddressHolder{institutionAddrres.AddressTypeId, institutionAddrres.InstitutionId, myInstitution.Name, institutionAddrres.Address, institutionAddrres.PostalCode}
				institutionsAddressHolder = append(institutionsAddressHolder, institutionAddress)
			}
		}
	}
	return institutionsAddressHolder
}

type InstitutionCourseHolder struct {
	InstitutionId     string
	CourseId          string
	InstitutionNane   string
	CourseName        string
	CourseDescription string
}

func GetInstitutionCourse() []InstitutionCourseHolder {
	var institutionsCourseHolder []InstitutionCourseHolder
	allInstitutionCourse, err := institutionIO.GetAllInstitutionCourses()
	if err != nil {
		fmt.Println("error reading allInstitutionCourse in GetInstitutionCourse method")
	} else {
		for _, institutionCourse := range allInstitutionCourse {
			institution, err := institutionIO.GetInstitution(institutionCourse.InstitutionId)
			if err != nil {
				fmt.Println("error reading institution in InstitutionManagementHandler method")
			}
			couse, err := academics.GetCourse(institutionCourse.CourseId)
			if err != nil {
				fmt.Println("error reading institution in InstitutionManagementHandler method")
			}
			if institution.Name != "" || couse.CourseName != "" {
				myInstitutionCours := InstitutionCourseHolder{institutionCourse.InstitutionId, institutionCourse.CourseId, institution.Name, couse.CourseName, couse.CourseDesc}
				institutionsCourseHolder = append(institutionsCourseHolder, myInstitutionCours)
			}
		}
	}
	return institutionsCourseHolder
}
