require 'rake'
require 'semver'

task :build do
    buildVersion = SemVer.find.to_s
    buildTimestamp = DateTime.now().strftime("%F %T")

    ldBuildVersion = "-X \"main.buildVersion=#{buildVersion}\""
    ldBuildTimestamp = "-X \"main.buildTimestamp=#{buildTimestamp}\""

    Dir.chdir('cmd') do
        sh('go', 'build', '-ldflags', "#{ldBuildVersion} #{ldBuildTimestamp}", '-o', 'dynapi.exe')
    end
end

task :test do
    sh("go test -v")
end

task :cover do
    sh("go test -coverprofile=coverage")
    sh("go tool cover -html=coverage")
end