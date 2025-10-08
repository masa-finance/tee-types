package industries

import "github.com/masa-finance/tee-types/pkg/util"

// IndustryId represents a LinkedIn industry identifier
type Id string

// Industry constants
const (
	AccommodationServices                               Id = "2190"
	FoodAndBeverageServices                             Id = "34"
	BarsTavernsAndNightclubs                            Id = "2217"
	Caterers                                            Id = "2212"
	MobileFoodServices                                  Id = "2214"
	Restaurants                                         Id = "32"
	Hospitality                                         Id = "31"
	BedAndBreakfastsHostelsHomestays                    Id = "2197"
	HotelsAndMotels                                     Id = "2194"
	AdministrativeAndSupportServices                    Id = "1912"
	CollectionAgencies                                  Id = "1938"
	EventsServices                                      Id = "110"
	FacilitiesServices                                  Id = "122"
	JanitorialServices                                  Id = "1965"
	LandscapingServices                                 Id = "2934"
	Fundraising                                         Id = "101"
	OfficeAdministration                                Id = "1916"
	SecurityAndInvestigations                           Id = "121"
	SecurityGuardsAndPatrolServices                     Id = "1956"
	SecuritySystemsServices                             Id = "1958"
	StaffingAndRecruiting                               Id = "104"
	ExecutiveSearchServices                             Id = "1923"
	TemporaryHelpServices                               Id = "1925"
	TelephoneCallCenters                                Id = "1931"
	TranslationAndLocalization                          Id = "108"
	TravelArrangements                                  Id = "30"
	WritingAndEditing                                   Id = "103"
	Construction                                        Id = "48"
	BuildingConstruction                                Id = "406"
	NonresIdentialBuildingConstruction                  Id = "413"
	ResIdentialBuildingConstruction                     Id = "408"
	CivilEngineering                                    Id = "51"
	HighwayStreetAndBrIdgeConstruction                  Id = "431"
	SubdivisionOfLand                                   Id = "428"
	UtilitySystemConstruction                           Id = "419"
	SpecialtyTradeContractors                           Id = "435"
	BuildingEquipmentContractors                        Id = "453"
	BuildingFinishingContractors                        Id = "460"
	BuildingStructureAndExteriorContractors             Id = "436"
	ConsumerServices                                    Id = "91"
	CivicAndSocialOrganizations                         Id = "90"
	Associations                                        Id = "1909"
	PoliticalOrganizations                              Id = "107"
	ProfessionalOrganizations                           Id = "1911"
	HouseholdServices                                   Id = "2318"
	NonProfitOrganizations                              Id = "100"
	PersonalAndLaundryServices                          Id = "2258"
	LaundryAndDrycleaningServices                       Id = "2272"
	PersonalCareServices                                Id = "2259"
	PetServices                                         Id = "2282"
	PhilanthropicFundraisingServices                    Id = "131"
	ReligiousInstitutions                               Id = "89"
	RepairAndMaintenance                                Id = "2225"
	CommercialAndIndustrialMachineryMaintenance         Id = "2247"
	ElectronicAndPrecisionEquipmentMaintenance          Id = "2240"
	FootwearAndLeatherGoodsRepair                       Id = "2255"
	ReupholsteryAndFurnitureRepair                      Id = "2253"
	VehicleRepairAndMaintenance                         Id = "2226"
	Education                                           Id = "1999"
	ELearningProvIders                                  Id = "132"
	HigherEducation                                     Id = "68"
	PrimaryAndSecondaryEducation                        Id = "67"
	ProfessionalTrainingAndCoaching                     Id = "105"
	TechnicalAndVocationalTraining                      Id = "2018"
	CosmetologyAndBarberSchools                         Id = "2019"
	FineArtsSchools                                     Id = "2025"
	FlightTraining                                      Id = "2020"
	LanguageSchools                                     Id = "2029"
	SecretarialSchools                                  Id = "2012"
	SportsAndRecreationInstruction                      Id = "2027"
	EntertainmentProvIders                              Id = "28"
	ArtistsAndWriters                                   Id = "38"
	MuseumsHistoricalSitesAndZoos                       Id = "37"
	HistoricalSites                                     Id = "2161"
	Museums                                             Id = "2159"
	ZoosAndBotanicalGardens                             Id = "2163"
	Musicians                                           Id = "115"
	PerformingArtsAndSpectatorSports                    Id = "2130"
	CircusesAndMagicShows                               Id = "2139"
	DanceCompanies                                      Id = "2135"
	PerformingArts                                      Id = "39"
	SpectatorSports                                     Id = "33"
	Racetracks                                          Id = "2143"
	SportsTeamsAndClubs                                 Id = "2142"
	TheaterCompanies                                    Id = "2133"
	RecreationalFacilities                              Id = "40"
	AmusementParksAndArcades                            Id = "2167"
	GamblingFacilitiesAndCasinos                        Id = "29"
	GolfCoursesAndCountryClubs                          Id = "2179"
	SkiingFacilities                                    Id = "2181"
	WellnessAndFitnessServices                          Id = "124"
	FarmingRanchingForestry                             Id = "201"
	Farming                                             Id = "63"
	Horticulture                                        Id = "150"
	ForestryAndLogging                                  Id = "298"
	RanchingAndFisheries                                Id = "256"
	Fisheries                                           Id = "66"
	Ranching                                            Id = "64"
	FinancialServices                                   Id = "43"
	CapitalMarkets                                      Id = "129"
	InvestmentAdvice                                    Id = "1720"
	InvestmentBanking                                   Id = "45"
	InvestmentManagement                                Id = "46"
	SecuritiesAndCommodityExchanges                     Id = "1713"
	VentureCapitalAndPrivateEquityPrincipals            Id = "106"
	CreditIntermediation                                Id = "1673"
	Banking                                             Id = "41"
	InternationalTradeAndDevelopment                    Id = "141"
	LoanBrokers                                         Id = "1696"
	SavingsInstitutions                                 Id = "1678"
	FundsAndTrusts                                      Id = "1742"
	InsuranceAndEmployeeBenefitFunds                    Id = "1743"
	PensionFunds                                        Id = "1745"
	TrustsAndEstates                                    Id = "1750"
	Insurance                                           Id = "42"
	ClaimsAdjustingActuarialServices                    Id = "1738"
	InsuranceAgenciesAndBrokerages                      Id = "1737"
	InsuranceCarriers                                   Id = "1725"
	GovernmentAdministration                            Id = "75"
	AdministrationOfJustice                             Id = "73"
	CorrectionalInstitutions                            Id = "3068"
	CourtsOfLaw                                         Id = "3065"
	FireProtection                                      Id = "3070"
	LawEnforcement                                      Id = "77"
	PublicSafety                                        Id = "78"
	EconomicPrograms                                    Id = "2375"
	TransportationPrograms                              Id = "3085"
	UtilitiesAdministration                             Id = "3086"
	EnvironmentalQualityPrograms                        Id = "388"
	AirWaterAndWasteProgramManagement                   Id = "2366"
	ConservationPrograms                                Id = "2368"
	HealthAndHumanServices                              Id = "2353"
	EducationAdministrationPrograms                     Id = "69"
	PublicAssistancePrograms                            Id = "2360"
	PublicHealth                                        Id = "2358"
	HousingAndCommunityDevelopment                      Id = "2369"
	CommunityDevelopmentAndUrbanPlanning                Id = "2374"
	HousingPrograms                                     Id = "3081"
	MilitaryAndInternationalAffairs                     Id = "2391"
	ArmedForces                                         Id = "71"
	InternationalAffairs                                Id = "74"
	PublicPolicyOffices                                 Id = "79"
	ExecutiveOffices                                    Id = "76"
	LegislativeOffices                                  Id = "72"
	SpaceResearchAndTechnology                          Id = "3089"
	HoldingCompanies                                    Id = "1905"
	HospitalsAndHealthCare                              Id = "14"
	CommunityServices                                   Id = "2115"
	ServicesForTheElderlyAndDisabled                    Id = "2112"
	Hospitals                                           Id = "2081"
	IndivIdualAndFamilyServices                         Id = "88"
	ChildDayCareServices                                Id = "2128"
	EmergencyAndReliefServices                          Id = "2122"
	VocationalRehabilitationServices                    Id = "2125"
	MedicalPractices                                    Id = "13"
	AlternativeMedicine                                 Id = "125"
	AmbulanceServices                                   Id = "2077"
	Chiropractors                                       Id = "2048"
	Dentists                                            Id = "2045"
	FamilyPlanningCenters                               Id = "2060"
	HomeHealthCareServices                              Id = "2074"
	MedicalAndDiagnosticLaboratories                    Id = "2069"
	MentalHealthCare                                    Id = "139"
	Optometrists                                        Id = "2050"
	OutpatientCareCenters                               Id = "2063"
	PhysicalOccupationalAndSpeechTherapists             Id = "2054"
	Physicians                                          Id = "2040"
	NursingHomesAndResIdentialCareFacilities            Id = "2091"
	Manufacturing                                       Id = "25"
	ApparelManufacturing                                Id = "598"
	FashionAccessoriesManufacturing                     Id = "615"
	AppliancesElectricalAndElectronicsManufacturing     Id = "112"
	ElectricLightingEquipmentManufacturing              Id = "998"
	ElectricalEquipmentManufacturing                    Id = "2468"
	FuelCellManufacturing                               Id = "3255"
	HouseholdApplianceManufacturing                     Id = "1005"
	ChemicalManufacturing                               Id = "54"
	AgriculturalChemicalManufacturing                   Id = "709"
	ArtificialRubberAndSyntheticFiberManufacturing      Id = "703"
	ChemicalRawMaterialsManufacturing                   Id = "690"
	PaintCoatingAndAdhesiveManufacturing                Id = "722"
	PersonalCareProductManufacturing                    Id = "18"
	PharmaceuticalManufacturing                         Id = "15"
	SoapAndCleaningProductManufacturing                 Id = "727"
	ClimateTechnologyProductManufacturing               Id = "3251"
	ComputersAndElectronicsManufacturing                Id = "24"
	AudioAndVIdeoEquipmentManufacturing                 Id = "973"
	CommunicationsEquipmentManufacturing                Id = "964"
	ComputerHardwareManufacturing                       Id = "3"
	AccessibleHardwareManufacturing                     Id = "3245"
	MagneticAndOpticalMediaManufacturing                Id = "994"
	MeasuringAndControlInstrumentManufacturing          Id = "983"
	SmartMeterManufacturing                             Id = "3254"
	SemiconductorManufacturing                          Id = "7"
	RenewableEnergySemiconductorManufacturing           Id = "144"
	FabricatedMetalProducts                             Id = "840"
	ArchitecturalAndStructuralMetalManufacturing        Id = "852"
	BoilersTanksAndShippingContainerManufacturing       Id = "861"
	ConstructionHardwareManufacturing                   Id = "871"
	CutleryAndHandtoolManufacturing                     Id = "849"
	MetalTreatments                                     Id = "883"
	MetalValveBallAndRollerManufacturing                Id = "887"
	SpringAndWireProductManufacturing                   Id = "873"
	TurnedProductsAndFastenerManufacturing              Id = "876"
	FoodAndBeverageManufacturing                        Id = "23"
	Breweries                                           Id = "562"
	Distilleries                                        Id = "564"
	Wineries                                            Id = "2500"
	AnimalFeedManufacturing                             Id = "481"
	BakedGoodsManufacturing                             Id = "529"
	BeverageManufacturing                               Id = "142"
	DairyProductManufacturing                           Id = "65"
	FruitAndVegetablePreservesManufacturing             Id = "504"
	MeatProductsManufacturing                           Id = "521"
	SeafoodProductManufacturing                         Id = "528"
	SugarAndConfectioneryProductManufacturing           Id = "495"
	FurnitureAndHomeFurnishingsManufacturing            Id = "26"
	HouseholdAndInstitutionalFurnitureManufacturing     Id = "1080"
	MattressAndBlindsManufacturing                      Id = "1095"
	OfficeFurnitureAndFixturesManufacturing             Id = "1090"
	GlassCeramicsAndConcreteManufacturing               Id = "145"
	AbrasivesAndNonmetallicMineralsManufacturing        Id = "799"
	ClayAndRefractoryProductsManufacturing              Id = "773"
	GlassProductManufacturing                           Id = "779"
	LimeAndGypsumProductsManufacturing                  Id = "794"
	LeatherProductManufacturing                         Id = "616"
	FootwearManufacturing                               Id = "622"
	WomensHandbagManufacturing                          Id = "625"
	MachineryManufacturing                              Id = "55"
	AgricultureConstructionMiningMachineryManufacturing Id = "901"
	AutomationMachineryManufacturing                    Id = "147"
	RobotManufacturing                                  Id = "3247"
	CommercialAndServiceIndustryMachineryManufacturing  Id = "918"
	EnginesAndPowerTransmissionEquipmentManufacturing   Id = "935"
	RenewableEnergyEquipmentManufacturing               Id = "3241"
	HvacAndRefrigerationEquipmentManufacturing          Id = "923"
	IndustrialMachineryManufacturing                    Id = "135"
	MetalworkingMachineryManufacturing                  Id = "928"
	MedicalEquipmentManufacturing                       Id = "17"
	OilAndCoalProductManufacturing                      Id = "679"
	PaperAndForestProductManufacturing                  Id = "61"
	PlasticsAndRubberProductManufacturing               Id = "743"
	PackagingAndContainersManufacturing                 Id = "146"
	PlasticsManufacturing                               Id = "117"
	RubberProductsManufacturing                         Id = "763"
	PrimaryMetalManufacturing                           Id = "807"
	PrintingServices                                    Id = "83"
	SportingGoodsManufacturing                          Id = "20"
	TextileManufacturing                                Id = "60"
	TobaccoManufacturing                                Id = "21"
	TransportationEquipmentManufacturing                Id = "1029"
	AviationAndAerospaceComponentManufacturing          Id = "52"
	DefenseAndSpaceManufacturing                        Id = "1"
	MotorVehicleManufacturing                           Id = "53"
	AlternativeFuelVehicleManufacturing                 Id = "3253"
	MotorVehiclePartsManufacturing                      Id = "1042"
	RailroadEquipmentManufacturing                      Id = "62"
	Shipbuilding                                        Id = "58"
	WoodProductManufacturing                            Id = "784"
	OilGasAndMining                                     Id = "332"
	Mining                                              Id = "56"
	CoalMining                                          Id = "341"
	MetalOreMining                                      Id = "345"
	NonmetallicMineralMining                            Id = "356"
	OilAndGas                                           Id = "57"
	NaturalGasExtraction                                Id = "3096"
	OilExtraction                                       Id = "3095"
	ProfessionalServices                                Id = "1810"
	Accounting                                          Id = "47"
	AdvertisingServices                                 Id = "80"
	GovernmentRelationsServices                         Id = "148"
	PublicRelationsAndCommunicationsServices            Id = "98"
	MarketResearch                                      Id = "97"
	ArchitectureAndPlanning                             Id = "50"
	AccessibleArchitectureAndDesign                     Id = "3246"
	BusinessConsultingAndServices                       Id = "11"
	EnvironmentalServices                               Id = "86"
	HumanResourcesServices                              Id = "137"
	MarketingServices                                   Id = "1862"
	OperationsConsulting                                Id = "2401"
	OutsourcingAndOffshoringConsulting                  Id = "123"
	StrategicManagementServices                         Id = "102"
	DesignServices                                      Id = "99"
	GraphicDesign                                       Id = "140"
	RegenerativeDesign                                  Id = "3256"
	InteriorDesign                                      Id = "3126"
	EngineeringServices                                 Id = "3242"
	RoboticsEngineering                                 Id = "3248"
	SurveyingAndMappingServices                         Id = "3249"
	ItServicesAndItConsulting                           Id = "96"
	ComputerAndNetworkSecurity                          Id = "118"
	DigitalAccessibilityServices                        Id = "3244"
	ItSystemCustomSoftwareDevelopment                   Id = "3102"
	ItSystemDataServices                                Id = "3106"
	ItSystemDesignServices                              Id = "1855"
	ItSystemInstallationAndDisposal                     Id = "3104"
	ItSystemOperationsAndMaintenance                    Id = "3103"
	ItSystemTestingAndEvaluation                        Id = "3107"
	ItSystemTrainingAndSupport                          Id = "3105"
	LegalServices                                       Id = "10"
	AlternativeDisputeResolution                        Id = "120"
	LawPractice                                         Id = "9"
	Photography                                         Id = "136"
	ResearchServices                                    Id = "70"
	BiotechnologyResearch                               Id = "12"
	NanotechnologyResearch                              Id = "114"
	ThinkTanks                                          Id = "130"
	ServicesForRenewableEnergy                          Id = "3243"
	VeterinaryServices                                  Id = "16"
	RealEstateAndEquipmentRentalServices                Id = "1757"
	EquipmentRentalServices                             Id = "1779"
	CommercialAndIndustrialEquipmentRental              Id = "1798"
	ConsumerGoodsRental                                 Id = "1786"
	RealEstate                                          Id = "44"
	LeasingNonResIdentialRealEstate                     Id = "128"
	LeasingResIdentialRealEstate                        Id = "1759"
	RealEstateAgentsAndBrokers                          Id = "1770"
	Retail                                              Id = "27"
	FoodAndBeverageRetail                               Id = "1339"
	RetailGroceries                                     Id = "22"
	OnlineAndMailOrderRetail                            Id = "1445"
	RetailApparelAndFashion                             Id = "19"
	RetailAppliancesElectricalAndElectronicEquipment    Id = "1319"
	RetailArtDealers                                    Id = "3186"
	RetailArtSupplies                                   Id = "111"
	RetailBooksAndPrintedNews                           Id = "1409"
	RetailBuildingMaterialsAndGardenEquipment           Id = "1324"
	RetailFlorists                                      Id = "1423"
	RetailFurnitureAndHomeFurnishings                   Id = "1309"
	RetailGasoline                                      Id = "1370"
	RetailHealthAndPersonalCareProducts                 Id = "1359"
	RetailPharmacies                                    Id = "3250"
	RetailLuxuryGoodsAndJewelry                         Id = "143"
	RetailMotorVehicles                                 Id = "1292"
	RetailMusicalInstruments                            Id = "1407"
	RetailOfficeEquipment                               Id = "138"
	RetailOfficeSuppliesAndGifts                        Id = "1424"
	RetailRecyclableMaterialsUsedMerchandise            Id = "1431"
	TechnologyInformationAndMedia                       Id = "1594"
	MediaTelecommunications                             Id = "3133"
	BookAndPeriodicalPublishing                         Id = "82"
	BookPublishing                                      Id = "1602"
	NewspaperPublishing                                 Id = "81"
	PeriodicalPublishing                                Id = "1600"
	BroadcastMediaProductionAndDistribution             Id = "36"
	CableAndSatelliteProgramming                        Id = "1641"
	RadioAndTelevisionBroadcasting                      Id = "1633"
	MoviesVIdeosAndSound                                Id = "35"
	AnimationAndPostProduction                          Id = "127"
	MediaProduction                                     Id = "126"
	MoviesAndSoundRecording                             Id = "1611"
	SoundRecording                                      Id = "1623"
	SheetMusicPublishing                                Id = "1625"
	Telecommunications                                  Id = "8"
	SatelliteTelecommunications                         Id = "1649"
	TelecommunicationsCarriers                          Id = "1644"
	WirelessServices                                    Id = "119"
	TechnologyInformationAndInternet                    Id = "6"
	DataInfrastructureAndAnalytics                      Id = "2458"
	BlockchainServices                                  Id = "3134"
	BusinessIntelligencePlatforms                       Id = "3128"
	ClimateDataAndAnalytics                             Id = "3252"
	InformationServices                                 Id = "84"
	InternetPublishing                                  Id = "3132"
	BusinessContent                                     Id = "3129"
	OnlineAudioAndVIdeoMedia                            Id = "113"
	InternetNews                                        Id = "3124"
	Libraries                                           Id = "85"
	Blogs                                               Id = "3125"
	InternetMarketplacePlatforms                        Id = "1285"
	SocialNetworkingPlatforms                           Id = "3127"
	SoftwareDevelopment                                 Id = "4"
	ComputerGames                                       Id = "109"
	MobileGamingApps                                    Id = "3131"
	ComputerNetworkingProducts                          Id = "5"
	DataSecuritySoftwareProducts                        Id = "3130"
	DesktopComputingSoftwareProducts                    Id = "3101"
	EmbeddedSoftwareProducts                            Id = "3099"
	MobileComputingSoftwareProducts                     Id = "3100"
	TransportationLogisticsSupplyChainAndStorage        Id = "116"
	AirlinesAndAviation                                 Id = "94"
	FreightAndPackageTransportation                     Id = "87"
	GroundPassengerTransportation                       Id = "1495"
	InterurbanAndRuralBusServices                       Id = "1504"
	SchoolAndEmployeeBusServices                        Id = "1512"
	ShuttlesAndSpecialNeedsTransportationServices       Id = "1517"
	SightseeingTransportation                           Id = "1532"
	TaxiAndLimousineServices                            Id = "1505"
	UrbanTransitServices                                Id = "1497"
	MaritimeTransportation                              Id = "95"
	PipelineTransportation                              Id = "1520"
	PostalServices                                      Id = "1573"
	RailTransportation                                  Id = "1481"
	TruckTransportation                                 Id = "92"
	WarehousingAndStorage                               Id = "93"
	Utilities                                           Id = "59"
	ElectricPowerGeneration                             Id = "383"
	FossilFuelElectricPowerGeneration                   Id = "385"
	NuclearElectricPowerGeneration                      Id = "386"
	RenewableEnergyPowerGeneration                      Id = "3240"
	BiomassElectricPowerGeneration                      Id = "390"
	GeothermalElectricPowerGeneration                   Id = "389"
	HydroelectricPowerGeneration                        Id = "384"
	SolarElectricPowerGeneration                        Id = "387"
	WindElectricPowerGeneration                         Id = "2489"
	ElectricPowerTransmissionControlAndDistribution     Id = "382"
	NaturalGasDistribution                              Id = "397"
	WaterWasteSteamAndAirConditioningServices           Id = "398"
	SteamAndAirConditioningSupply                       Id = "404"
	WasteCollection                                     Id = "1981"
	WasteTreatmentAndDisposal                           Id = "1986"
	WaterSupplyAndIrrigationSystems                     Id = "400"
	Wholesale                                           Id = "133"
	WholesaleAlcoholicBeverages                         Id = "1267"
	WholesaleApparelAndSewingSupplies                   Id = "1222"
	WholesaleAppliancesElectricalAndElectronics         Id = "1171"
	WholesaleBuildingMaterials                          Id = "49"
	WholesaleChemicalAndAlliedProducts                  Id = "1257"
	WholesaleComputerEquipment                          Id = "1157"
	WholesaleDrugsAndSundries                           Id = "1221"
	WholesaleFoodAndBeverage                            Id = "1231"
	WholesaleFootwear                                   Id = "1230"
	WholesaleFurnitureAndHomeFurnishings                Id = "1137"
	WholesaleHardwarePlumbingHeatingEquipment           Id = "1178"
	WholesaleImportAndExport                            Id = "134"
	WholesaleLuxuryGoodsAndJewelry                      Id = "1208"
	WholesaleMachinery                                  Id = "1187"
	WholesaleMetalsAndMinerals                          Id = "1166"
	WholesaleMotorVehiclesAndParts                      Id = "1128"
	WholesalePaperProducts                              Id = "1212"
	WholesalePetroleumAndPetroleumProducts              Id = "1262"
	WholesalePhotographyEquipmentAndSupplies            Id = "1153"
	WholesaleRawFarmProducts                            Id = "1250"
	WholesaleRecyclableMaterials                        Id = "1206"
)

var All = util.NewSet(
	AccommodationServices,
	FoodAndBeverageServices,
	BarsTavernsAndNightclubs,
	Caterers,
	MobileFoodServices,
	Restaurants,
	Hospitality,
	BedAndBreakfastsHostelsHomestays,
	HotelsAndMotels,
	AdministrativeAndSupportServices,
	CollectionAgencies,
	EventsServices,
	FacilitiesServices,
	JanitorialServices,
	LandscapingServices,
	Fundraising,
	OfficeAdministration,
	SecurityAndInvestigations,
	SecurityGuardsAndPatrolServices,
	SecuritySystemsServices,
	StaffingAndRecruiting,
	ExecutiveSearchServices,
	TemporaryHelpServices,
	TelephoneCallCenters,
	TranslationAndLocalization,
	TravelArrangements,
	WritingAndEditing,
	Construction,
	BuildingConstruction,
	CivilEngineering,
	SubdivisionOfLand,
	UtilitySystemConstruction,
	SpecialtyTradeContractors,
	BuildingEquipmentContractors,
	BuildingFinishingContractors,
	BuildingStructureAndExteriorContractors,
	ConsumerServices,
	CivicAndSocialOrganizations,
	Associations,
	PoliticalOrganizations,
	ProfessionalOrganizations,
	HouseholdServices,
	NonProfitOrganizations,
	PersonalAndLaundryServices,
	LaundryAndDrycleaningServices,
	PersonalCareServices,
	PetServices,
	PhilanthropicFundraisingServices,
	ReligiousInstitutions,
	RepairAndMaintenance,
	CommercialAndIndustrialMachineryMaintenance,
	ElectronicAndPrecisionEquipmentMaintenance,
	FootwearAndLeatherGoodsRepair,
	ReupholsteryAndFurnitureRepair,
	VehicleRepairAndMaintenance,
	Education,
	HigherEducation,
	PrimaryAndSecondaryEducation,
	ProfessionalTrainingAndCoaching,
	TechnicalAndVocationalTraining,
	CosmetologyAndBarberSchools,
	FineArtsSchools,
	FlightTraining,
	LanguageSchools,
	SecretarialSchools,
	SportsAndRecreationInstruction,
	ArtistsAndWriters,
	MuseumsHistoricalSitesAndZoos,
	HistoricalSites,
	Museums,
	ZoosAndBotanicalGardens,
	Musicians,
	PerformingArtsAndSpectatorSports,
	CircusesAndMagicShows,
	DanceCompanies,
	PerformingArts,
	SpectatorSports,
	Racetracks,
	SportsTeamsAndClubs,
	TheaterCompanies,
	RecreationalFacilities,
	AmusementParksAndArcades,
	GamblingFacilitiesAndCasinos,
	GolfCoursesAndCountryClubs,
	SkiingFacilities,
	WellnessAndFitnessServices,
	FarmingRanchingForestry,
	Farming,
	Horticulture,
	ForestryAndLogging,
	RanchingAndFisheries,
	Fisheries,
	Ranching,
	FinancialServices,
	CapitalMarkets,
	InvestmentAdvice,
	InvestmentBanking,
	InvestmentManagement,
	SecuritiesAndCommodityExchanges,
	VentureCapitalAndPrivateEquityPrincipals,
	CreditIntermediation,
	Banking,
	InternationalTradeAndDevelopment,
	LoanBrokers,
	SavingsInstitutions,
	FundsAndTrusts,
	InsuranceAndEmployeeBenefitFunds,
	PensionFunds,
	TrustsAndEstates,
	Insurance,
	ClaimsAdjustingActuarialServices,
	InsuranceAgenciesAndBrokerages,
	InsuranceCarriers,
	GovernmentAdministration,
	AdministrationOfJustice,
	CorrectionalInstitutions,
	CourtsOfLaw,
	FireProtection,
	LawEnforcement,
	PublicSafety,
	EconomicPrograms,
	TransportationPrograms,
	UtilitiesAdministration,
	EnvironmentalQualityPrograms,
	AirWaterAndWasteProgramManagement,
	ConservationPrograms,
	HealthAndHumanServices,
	EducationAdministrationPrograms,
	PublicAssistancePrograms,
	PublicHealth,
	HousingAndCommunityDevelopment,
	CommunityDevelopmentAndUrbanPlanning,
	HousingPrograms,
	MilitaryAndInternationalAffairs,
	ArmedForces,
	InternationalAffairs,
	PublicPolicyOffices,
	ExecutiveOffices,
	LegislativeOffices,
	SpaceResearchAndTechnology,
	HoldingCompanies,
	HospitalsAndHealthCare,
	CommunityServices,
	ServicesForTheElderlyAndDisabled,
	Hospitals,
	ChildDayCareServices,
	EmergencyAndReliefServices,
	VocationalRehabilitationServices,
	MedicalPractices,
	AlternativeMedicine,
	AmbulanceServices,
	Chiropractors,
	Dentists,
	FamilyPlanningCenters,
	HomeHealthCareServices,
	MedicalAndDiagnosticLaboratories,
	MentalHealthCare,
	Optometrists,
	OutpatientCareCenters,
	PhysicalOccupationalAndSpeechTherapists,
	Physicians,
	Manufacturing,
	ApparelManufacturing,
	FashionAccessoriesManufacturing,
	AppliancesElectricalAndElectronicsManufacturing,
	ElectricLightingEquipmentManufacturing,
	ElectricalEquipmentManufacturing,
	FuelCellManufacturing,
	HouseholdApplianceManufacturing,
	ChemicalManufacturing,
	AgriculturalChemicalManufacturing,
	ArtificialRubberAndSyntheticFiberManufacturing,
	ChemicalRawMaterialsManufacturing,
	PaintCoatingAndAdhesiveManufacturing,
	PersonalCareProductManufacturing,
	PharmaceuticalManufacturing,
	SoapAndCleaningProductManufacturing,
	ClimateTechnologyProductManufacturing,
	ComputersAndElectronicsManufacturing,
	CommunicationsEquipmentManufacturing,
	ComputerHardwareManufacturing,
	AccessibleHardwareManufacturing,
	MagneticAndOpticalMediaManufacturing,
	MeasuringAndControlInstrumentManufacturing,
	SmartMeterManufacturing,
	SemiconductorManufacturing,
	RenewableEnergySemiconductorManufacturing,
	FabricatedMetalProducts,
	ArchitecturalAndStructuralMetalManufacturing,
	BoilersTanksAndShippingContainerManufacturing,
	ConstructionHardwareManufacturing,
	CutleryAndHandtoolManufacturing,
	MetalTreatments,
	MetalValveBallAndRollerManufacturing,
	SpringAndWireProductManufacturing,
	TurnedProductsAndFastenerManufacturing,
	FoodAndBeverageManufacturing,
	Breweries,
	Distilleries,
	Wineries,
	AnimalFeedManufacturing,
	BakedGoodsManufacturing,
	BeverageManufacturing,
	DairyProductManufacturing,
	FruitAndVegetablePreservesManufacturing,
	MeatProductsManufacturing,
	SeafoodProductManufacturing,
	SugarAndConfectioneryProductManufacturing,
	FurnitureAndHomeFurnishingsManufacturing,
	HouseholdAndInstitutionalFurnitureManufacturing,
	MattressAndBlindsManufacturing,
	OfficeFurnitureAndFixturesManufacturing,
	GlassCeramicsAndConcreteManufacturing,
	AbrasivesAndNonmetallicMineralsManufacturing,
	ClayAndRefractoryProductsManufacturing,
	GlassProductManufacturing,
	LimeAndGypsumProductsManufacturing,
	LeatherProductManufacturing,
	FootwearManufacturing,
	WomensHandbagManufacturing,
	MachineryManufacturing,
	AgricultureConstructionMiningMachineryManufacturing,
	AutomationMachineryManufacturing,
	RobotManufacturing,
	CommercialAndServiceIndustryMachineryManufacturing,
	EnginesAndPowerTransmissionEquipmentManufacturing,
	RenewableEnergyEquipmentManufacturing,
	HvacAndRefrigerationEquipmentManufacturing,
	IndustrialMachineryManufacturing,
	MetalworkingMachineryManufacturing,
	MedicalEquipmentManufacturing,
	OilAndCoalProductManufacturing,
	PaperAndForestProductManufacturing,
	PlasticsAndRubberProductManufacturing,
	PackagingAndContainersManufacturing,
	PlasticsManufacturing,
	RubberProductsManufacturing,
	PrimaryMetalManufacturing,
	PrintingServices,
	SportingGoodsManufacturing,
	TextileManufacturing,
	TobaccoManufacturing,
	TransportationEquipmentManufacturing,
	AviationAndAerospaceComponentManufacturing,
	DefenseAndSpaceManufacturing,
	MotorVehicleManufacturing,
	AlternativeFuelVehicleManufacturing,
	MotorVehiclePartsManufacturing,
	RailroadEquipmentManufacturing,
	Shipbuilding,
	WoodProductManufacturing,
	OilGasAndMining,
	Mining,
	CoalMining,
	MetalOreMining,
	NonmetallicMineralMining,
	OilAndGas,
	NaturalGasExtraction,
	OilExtraction,
	ProfessionalServices,
	Accounting,
	AdvertisingServices,
	GovernmentRelationsServices,
	PublicRelationsAndCommunicationsServices,
	MarketResearch,
	ArchitectureAndPlanning,
	AccessibleArchitectureAndDesign,
	BusinessConsultingAndServices,
	EnvironmentalServices,
	HumanResourcesServices,
	MarketingServices,
	OperationsConsulting,
	OutsourcingAndOffshoringConsulting,
	StrategicManagementServices,
	DesignServices,
	GraphicDesign,
	RegenerativeDesign,
	InteriorDesign,
	EngineeringServices,
	RoboticsEngineering,
	SurveyingAndMappingServices,
	ItServicesAndItConsulting,
	ComputerAndNetworkSecurity,
	DigitalAccessibilityServices,
	ItSystemCustomSoftwareDevelopment,
	ItSystemDataServices,
	ItSystemDesignServices,
	ItSystemInstallationAndDisposal,
	ItSystemOperationsAndMaintenance,
	ItSystemTestingAndEvaluation,
	ItSystemTrainingAndSupport,
	LegalServices,
	AlternativeDisputeResolution,
	LawPractice,
	Photography,
	ResearchServices,
	BiotechnologyResearch,
	NanotechnologyResearch,
	ThinkTanks,
	ServicesForRenewableEnergy,
	VeterinaryServices,
	RealEstateAndEquipmentRentalServices,
	EquipmentRentalServices,
	CommercialAndIndustrialEquipmentRental,
	ConsumerGoodsRental,
	RealEstate,
	RealEstateAgentsAndBrokers,
	Retail,
	FoodAndBeverageRetail,
	RetailGroceries,
	OnlineAndMailOrderRetail,
	RetailApparelAndFashion,
	RetailAppliancesElectricalAndElectronicEquipment,
	RetailArtDealers,
	RetailArtSupplies,
	RetailBooksAndPrintedNews,
	RetailBuildingMaterialsAndGardenEquipment,
	RetailFlorists,
	RetailFurnitureAndHomeFurnishings,
	RetailGasoline,
	RetailHealthAndPersonalCareProducts,
	RetailPharmacies,
	RetailLuxuryGoodsAndJewelry,
	RetailMotorVehicles,
	RetailMusicalInstruments,
	RetailOfficeEquipment,
	RetailOfficeSuppliesAndGifts,
	RetailRecyclableMaterialsUsedMerchandise,
	TechnologyInformationAndMedia,
	MediaTelecommunications,
	BookAndPeriodicalPublishing,
	BookPublishing,
	NewspaperPublishing,
	PeriodicalPublishing,
	BroadcastMediaProductionAndDistribution,
	CableAndSatelliteProgramming,
	RadioAndTelevisionBroadcasting,
	AnimationAndPostProduction,
	MediaProduction,
	MoviesAndSoundRecording,
	SoundRecording,
	SheetMusicPublishing,
	Telecommunications,
	SatelliteTelecommunications,
	TelecommunicationsCarriers,
	WirelessServices,
	TechnologyInformationAndInternet,
	DataInfrastructureAndAnalytics,
	BlockchainServices,
	BusinessIntelligencePlatforms,
	ClimateDataAndAnalytics,
	InformationServices,
	InternetPublishing,
	BusinessContent,
	InternetNews,
	Libraries,
	Blogs,
	InternetMarketplacePlatforms,
	SocialNetworkingPlatforms,
	SoftwareDevelopment,
	ComputerGames,
	MobileGamingApps,
	ComputerNetworkingProducts,
	DataSecuritySoftwareProducts,
	DesktopComputingSoftwareProducts,
	EmbeddedSoftwareProducts,
	MobileComputingSoftwareProducts,
	TransportationLogisticsSupplyChainAndStorage,
	AirlinesAndAviation,
	FreightAndPackageTransportation,
	GroundPassengerTransportation,
	InterurbanAndRuralBusServices,
	SchoolAndEmployeeBusServices,
	ShuttlesAndSpecialNeedsTransportationServices,
	SightseeingTransportation,
	TaxiAndLimousineServices,
	UrbanTransitServices,
	MaritimeTransportation,
	PipelineTransportation,
	PostalServices,
	RailTransportation,
	TruckTransportation,
	WarehousingAndStorage,
	Utilities,
	ElectricPowerGeneration,
	FossilFuelElectricPowerGeneration,
	NuclearElectricPowerGeneration,
	RenewableEnergyPowerGeneration,
	BiomassElectricPowerGeneration,
	GeothermalElectricPowerGeneration,
	HydroelectricPowerGeneration,
	SolarElectricPowerGeneration,
	WindElectricPowerGeneration,
	ElectricPowerTransmissionControlAndDistribution,
	NaturalGasDistribution,
	WaterWasteSteamAndAirConditioningServices,
	SteamAndAirConditioningSupply,
	WasteCollection,
	WasteTreatmentAndDisposal,
	WaterSupplyAndIrrigationSystems,
	Wholesale,
	WholesaleAlcoholicBeverages,
	WholesaleApparelAndSewingSupplies,
	WholesaleAppliancesElectricalAndElectronics,
	WholesaleBuildingMaterials,
	WholesaleChemicalAndAlliedProducts,
	WholesaleComputerEquipment,
	WholesaleDrugsAndSundries,
	WholesaleFoodAndBeverage,
	WholesaleFootwear,
	WholesaleFurnitureAndHomeFurnishings,
	WholesaleHardwarePlumbingHeatingEquipment,
	WholesaleImportAndExport,
	WholesaleLuxuryGoodsAndJewelry,
	WholesaleMachinery,
	WholesaleMetalsAndMinerals,
	WholesaleMotorVehiclesAndParts,
	WholesalePaperProducts,
	WholesalePetroleumAndPetroleumProducts,
	WholesalePhotographyEquipmentAndSupplies,
	WholesaleRawFarmProducts,
	WholesaleRecyclableMaterials,
)
